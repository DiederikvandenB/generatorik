build:
	go build && go install

copy-templates:
	mkdir -p ~/.config/generatorik/
	cp -R ./template-examples/* ~/.config/generatorik/templates
