VERSION:=0.3.4
IMAGE:=cr.yandex/crpkmcbem8um7rd1gk5i/hoba-hoba-bot

build:
	docker build . -t ${IMAGE}:${VERSION} -t ${IMAGE}

push:
	docker push ${IMAGE}:${VERSION}
	docker push ${IMAGE}

deploy: build push

# сборка под x86 на ARM
deploy-m1:
	docker buildx build --platform linux/amd64 -t ${IMAGE}:${VERSION} -t ${IMAGE} --push .
