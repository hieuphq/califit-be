PROJECT_DIR=github.com/hieuphq/califit-be
PROJECT_ROOT=${GOPATH}/src/${PROJECT_DIR}
PROJECT_NAME=califit-be

#build alpine
build-alpine:
	CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o bin/${PROJECT_NAME} ./ext/cmd/${PROJECT_NAME}

# build
cln-bin:
	rm -rf bin
bin-migrator:
	go build -o bin/migrator ./ext/cmd/migrator
bin-${PROJECT_NAME}:
	go build -i -o bin/${PROJECT_NAME} ./ext/cmd/${PROJECT_NAME}
bin: cln-bin bin-migrator

# make local dev environment
local-db: bin
	eval "docker-compose -f localdb-docker-compose.yaml down"
	eval "docker-compose -f localdb-docker-compose.yaml up -d"

env-file:
	@if ! [ -e ".env_migrator.yaml" ] ; then cat .env_migrator.yaml.example > .env_migrator.yaml; fi
	@if ! [ -e "sqlboiler.toml" ] ; then cat sqlboiler.toml.example > sqlboiler.toml; fi
	@if ! [ -e ".env" ] ; then cat .env.example > .env ; fi

local-env: local-db env-file
	@echo "Waiting for database connection..."
	@while ! docker exec backend_${PROJECT_NAME} pg_isready -h localhost -p 5432 > /dev/null; do \
		sleep 1; \
	done
	bin/migrator up	

# dep is dependencies tools
depinit:
	cd ext/cmd/;dep init
dep:
	cd ext/cmd/;dep ensure -update
depcln:
	cd ext/cmd/;rm -rf vendor lock.json manifest.json

# gen code
# GOA
gen-goa-app:
	cd goa/;goagen app -d ${PROJECT_DIR}/goa/design
gen-goa-client:
	cd goa/;goagen client -d ${PROJECT_DIR}/goa/design
gen-goa-swagger:
	goagen swagger -d ${PROJECT_DIR}/goa/design
gen-goa-controller:
	cd ext/controller/;goagen controller -d ${PROJECT_DIR}/goa/design
gen-goa: gen-goa-app gen-goa-client gen-goa-swagger gen-goa-controller

# SQL_BOILER
gen-sqlboiler-model:
	@sqlboiler --wipe psql
	
gen-workaround-vendor-goa-before:
	mv vendor vendor_wk
gen-workaround-vendor-goa-after:
	mv vendor_wk vendor
gen: gen-sqlboiler-model gen-workaround-vendor-goa-before gen-goa gen-workaround-vendor-goa-after

# Test
test:
	go test ./...

# Dev
dev: bin-${PROJECT_NAME}
	ENV=DEV bin/${PROJECT_NAME}

# RSA Key
clean-rsa-key:
	rm -rf rsa_key rsa_key.pub

rsa-key:clean-rsa-key
	@openssl genrsa -out rsa_key 2048
	@openssl rsa -in rsa_key -pubout > rsa_key.pub
