build-all:
	make -C frontend build
	mkdir -p server/golang/frontend
	cp -r frontend/build/* server/golang/frontend/
	make -C server/golang build

remove-binary:
	sudo rm -f /usr/local/bin/sb

install: build-all remove-binary
	sudo cp server/golang/bin/snmp-browser /usr/local/bin/.

install-short-name: build-all remove-binary
	sudo cp server/golang/bin/snmp-browser /usr/local/bin/sb

build-all-docker: clean
	docker build --tag hegemonies/snmp-browser:latest .

clean:
	rm -rf frontend/build
	rm -rf server/golang/bin
	rm -rf server/golang/frontend

push-docker-image:
	docker push hegemonies/snmp-browser:latest
