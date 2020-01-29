# Starts the datastore emulator for running locally. Called by `make test`.
datastore-start:
	@gcloud beta emulators datastore start --no-store-on-disk --host-port=127.0.0.1:8084 --consistency 1.0 --quiet > /dev/null 2>&1 &
	@echo "Cloud Datastore Emulator started..."

# Looks for a running datastore emulator and stops it.
datastore-stop:
	@kill -9 `ps ax | grep 'CloudDatastore.jar' | grep -v grep | awk '{print $1}'` > /dev/null 2>&1 &
	@echo "Cloud Datastore Emulator stopped"

test: datastore-start
	@export DATASTORE_EMULATOR_HOST=127.0.0.1:8084; \
	go test -v ./...
	@$(MAKE) datastore-stop
