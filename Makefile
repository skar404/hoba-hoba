VERSION:=0.0.1
IMAGE:=cr.yandex/crpkmcbem8um7rd1gk5i/hoba-hoba-bot

build:
	docker build . -t ${IMAGE}:${VERSION} -t ${IMAGE}

push:
	docker push ${IMAGE}:${VERSION}
	docker push ${IMAGE}

deploy: build push