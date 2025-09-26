# Собираем все цели из этого файла в качестве значения переменной ALL_TARGETS
ALL_TARGETS := $(shell grep -E '^[a-zA-Z0-9_-]+:' Makefile | cut -d: -f1 | grep -v '^\.' | sort -u)

# Убираем конфликты названий целей с одноимёнными файлами
.PHONY: $(ALL_TARGETS)

# Определяем последнюю по числу папку для добавления в репу, например L3.21
LAST_TASK := $(shell ls -d L*.* | sort -V | tail -n1)

# Первый аргумент вызова make сохраним в переменную PARAM
# Пример: 
# > make Arg1 target1 target2 ...
# Arg1 сохраняем в переменную PARAM:
PARAM := $(firstword $(MAKECMDGOALS))

##############################################################################
#   Ниже идёт блок-условие, который обрабатывает входной параметр PARAM,
#   в зависимосим от его содержания. Если PARAM - одна из целей текущего 
#   файла, то запускается эта цель, если нет, то предполагается, что PARAM
#   содержит название подпапки, в которой требуется запустить проект
#   Например, при вызове из консоли make L1.1, в цель run передастся 
#   значение подпапки L1.1, в которой будет запущен проект
#
ifeq ($(filter $(PARAM),$(ALL_TARGETS)),) # если PARAM есть в списке целей ALL_TARGETS, то функция возвращает PARAM, иначе возвращает пустоту
ifeq ($(PARAM),) # если PARAM пустой, то выходим из дальнейших условий и выполняем стандартный ход событий
else # если PARAM не пустой (то есть не является одной из целей в этом файле), то..
$(eval .PHONY: $(PARAM)) # realtime (параметр eval) добавляем его к списку названий всех целей для устранения конфликта при запуске на случай наличия одноименного файла
$(eval $(PARAM): ; @$(MAKE) run LAST_TASK=$(PARAM)) # создаём динамическую цель PARAM: и сразу запускаем её (она делает переадресацию на цель run) при этом переопределяя переменную LAST_TASK = PARAM
endif
endif

# Автоматическая цель для запуска проекта в папке LAST_TASK
run: init_new_task
	@cd $(LAST_TASK) && go run .

###############################################################################
##                                                                           ##
##                    CОЗДАНИЕ НОВОГО РЕПОЗИТОРИЯ WB_L..                     ##
##                                                                           ##
###############################################################################

# Обрезаем имя после точки для создания репозитория, например L3.21 до L3
CUTED_LAST_TASK := $(shell echo $(LAST_TASK) | cut -d. -f1)

# Имя репозитория, например WB_L3
REPO_NAME := WB_$(CUTED_LAST_TASK)

# Токен GitHub (лежит уровнем выше и не вставляется в репу для сохранения приватности репы)
ifeq ($(wildcard ../github_privacy/gh_tok.en),)
GITHUB_TOKEN :=
else
GITHUB_TOKEN := $(shell cat ../github_privacy/gh_tok.en)
endif

# URL API GitHub
GITHUB_API := https://api.github.com/user/repos

# URL репозитория
GIT_URL := git@github.com:mysokolsky/$(REPO_NAME).git

# Создаём репозиторий, если его ещё нет
create-repo:
	@if git ls-remote $(GIT_URL) >/dev/null 2>&1; then \
		echo "Репозиторий $(REPO_NAME) уже существует"; \
	else \
		echo "Создаём репозиторий $(REPO_NAME)"; \
		curl -s -H "Authorization: token $(GITHUB_TOKEN)" \
		     -H "Accept: application/vnd.github.v3+json" \
		     $(GITHUB_API) \
		     -d '{"name":"$(REPO_NAME)","private":false,"auto_init":true}'; \
	fi

###############################################################################
##                                                                           ##
##                    ЗАЛИВКА ПОСЛЕДНИХ ИЗМЕНЕНИЙ В РЕПУ                     ##
##                                                                           ##
###############################################################################

# Определяем ветку для автоматического пуша
GITBRANCH:=$(shell git rev-parse --abbrev-ref HEAD)

# Дальше цели для работы с гитом
add:
	@git add "$(LAST_TASK)"; sleep 1
	@for f in *; do \
		if [ -f "$$f" ]; then \
			git add "$$f" 2>/dev/null || true; \
		fi; \
	done

commit:
	@git commit -m "=== $(LAST_TASK) == $(GITBRANCH) === $(shell date +'ДАТА %d-%m-%y === ВРЕМЯ %H:%M:%S') ====="; sleep 1

push: add 
	@if git diff --cached --quiet; then \
		echo "Нет изменений для сохранения."; \
	else \
		$(MAKE) commit; \
		git push origin $(GITBRANCH); \
	fi

# автоматическая загрузка с гит-репозитория на текущую машину
pull:
	git stash
	git pull origin $(GITBRANCH)

###############################################################################
##                                                                           ##
##                              СЛУЖЕБНЫЕ ЦЕЛИ                               ##
##                                                                           ##
###############################################################################

# Сохранить токен для приватного доступа к гитхабу
gh_token_save:
	git config --global credential.helper osxkeychain

# Убрать предупрежение о включении игногируемых файлов в репу
skip_attention:
	git config advice.addIgnoredFile false

# инициализировать голанг-проект в папке последнего задания
init_new_task:
	@if [ ! -f $(LAST_TASK)/go.mod ]; then \
		echo "go.mod не найден, инициализируем новый проект..."; \
		echo; \
		cd $(LAST_TASK) && { \
			go mod init github.com/mysokolsky/$(REPO_NAME)/$(LAST_TASK) > /dev/null 2>&1; \
			go mod tidy  > /dev/null 2>&1; \
		}; \
	fi