include('./backend/publicapi/config/Tiltfile')
include('./backend/roomsocket/config/Tiltfile')
include('./frontend/config/Tiltfile')

load('ext://helm_resource', 'helm_resource', 'helm_repo')
helm_repo('bitnami', 'https://charts.bitnami.com/bitnami')
helm_resource('rabbitmq', 'bitnami/rabbitmq')
