# ansicix
A Simple Devops Tool for Kubernetes pods  HealthCheck created by golang Cobra 

````C
project-root-directory/
│
├── ansible/
│   ├── inventory/
│   │   ├── hosts.ini
│   │
│   ├── playbooks/
│   │   ├── main.yml (or kubernetes_health_check.yml)
│   │
│   └── ansible.cfg
│
├── python/
│   └── send_to_elasticsearch.py
│
├── go/
│   ├── ansicix.go
│   └── go.mod
│
├── kubernetes/
│   └── app-deployment.yml
│
└── README.md
````
