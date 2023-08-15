import requests
import sys

def send_to_elasticsearch(status):
    data = {
        "status": status
    }
    response = requests.post("http://YOUR_ELASTICSEARCH_URL:9200/k8s_health_status/_doc", json=data)
    if response.status_code != 201:
        print("Error sending data to Elasticsearch:", response.text)

if __name__ == "__main__":
    status = sys.argv[1]
    send_to_elasticsearch(status)

