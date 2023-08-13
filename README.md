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
│   └── go.mod (ve diğer Go projeleri için gereken dosyalar)
│
├── kubernetes/
│   └── app-deployment.yml
│
└── README.md (veya başka bir adla belgelendirme/markdown dosyası)
````
