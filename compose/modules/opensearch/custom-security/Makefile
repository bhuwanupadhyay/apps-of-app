initialize:
	./init-certs.sh && sudo ./init-vm.sh

build_docker:
	docker-compose build

start_docker: build_docker
	docker-compose up -d

clean_docker:
	docker-compose down -v --remove-orphans

remove_docker:
	docker-compose down --rmi all -v --remove-orphans

logs_node1:
	docker logs opensearch-cluster-node1 -f | tee node.log

logs_dashboard:
	docker logs opensearch-cluster-dashboard -f | tee dashboard.log