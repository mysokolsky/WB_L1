gh_token_save:
	git config --global credential.helper osxkeychain

push:
	git add .
	git commit -m "new"
	git push origin master