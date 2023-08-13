# ansicix
A Simple Devops Tool for Kubernetes pods  HealthCheck created by golang Cobra 

project-root-directory/
│
├── ansible/
│   ├── inventory/
│   │   ├── hosts.ini
│   │
│   ├── playbooks/
│   │   ├── main.yml (veya kubernetes_health_check.yml ya da benzeri bir isim)
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
