
ENVIRONMENT ?= development

define tailwind
.build/$1/public/css/tailwind.css: services/$1/css/tailwind.css services/$1/tailwind.config.js
	@echo Building Tailwind for service $1
	NODE_ENV=$(ENVIRONMENT) npx tailwindcss-cli@0.1.2 build services/$1/css/tailwind.css -c services/$1/tailwind.config.js -o .build/$1/public/css/tailwind.css

.PHONY: $1_tailwind
$1_tailwind: .build/$1/public/css/tailwind.css
endef

define views
.build/$1/views: services/$1/views/*.html
	rm -Rf .build/$1/views
	cp -Rf services/$1/views .build/$1/views

.PHONY: $1_views
$1_views: .build/$1/views
endef

define js
.build/$1/public/js/bundle.js: services/$1/js/*
	npx webpack --mode=development

.PHONY: $1_js
$1_js: .build/$1/public/js/bundle.js
endef