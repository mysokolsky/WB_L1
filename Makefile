# Определяем последнюю по числу папку для добавления в репу
LAST_TASK:=$(shell ls -d L*.* | sort -V | tail -n1)

# Определяем ветку для автоматического пуша
GITBRANCH:=$(shell git rev-parse --abbrev-ref HEAD)


########### Дальше цели для работы с гитом
add:
	@git add "$(LAST_TASK)"; sleep 1
	@for f in *; do \
		if [ -f "$$f" ]; then \
			git add "$$f" || true; \
		fi; \
	done

commit:
	@git commit -m "=== $(LAST_TASK) == $(GITBRANCH) === $(shell date +'ДАТА %d-%m-%y === ВРЕМЯ %H:%M:%S') ====="; sleep 1

push: add 
	@if git diff --cached --quiet; then \
		echo "Нет изменений для сохранения."; \
	else \
		@make commit; \
		git push origin $(GITBRANCH); \
	fi

# автоматическая загрузка с гит-репозитория на текущую машину
pull:
	git stash
	git pull origin $(GITBRANCH)


################## Приватный доступ к гитхабу по токену
gh_token_save:
	git config --global credential.helper osxkeychain