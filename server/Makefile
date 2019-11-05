CONTAINER_IMAGE=maestro-dev-server
PORT=8080

RED=\e[31m
RED_TERM=\e[m

dev: 
	docker build -t $(CONTAINER_IMAGE) .
	docker run --rm -p $(PORT):8080 -d --name $(CONTAINER_IMAGE) $(CONTAINER_IMAGE)
	@echo  "\n$(RED)maestro server is running on port:$(PORT)$(RED_TERM)"

clean-dev:
	docker container stop $(CONTAINER_IMAGE)

.PHONY: dev clean-dev
