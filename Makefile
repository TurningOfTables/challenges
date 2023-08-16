add:
	mkdir $(NAME)
	echo "package $(NAME)" > $(NAME)/$(NAME).go
	echo "package $(NAME)" > $(NAME)/$(NAME)_test.go
	cd $(NAME) && go mod init challenges/$(NAME)
	go work use $(NAME)