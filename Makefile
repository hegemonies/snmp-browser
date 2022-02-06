build-all:
	make -C frontend build
	mkdir -R server/golang/frontend
	cp -r frontend/build server/golang/frontend/
	make -C server/golang build

build-all-docker:
	docker build --tag hegemonies/snmp-browser:latest .

clean:
	rm -rf frontend/build
	rm -rf server/golang/bin
	rm -rf server/golang/frontend
