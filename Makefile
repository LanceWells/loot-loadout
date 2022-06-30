gen-api:
	cd api && \
	buf generate && \
	go generate ./api/...

gen-manifests:
	kubectl kustomize ./backend/publicapi/config/overlays/prod/ -o manifests/publicapi.yaml && \
	kubectl kustomize ./backend/roomsocket/config/overlays/prod/ -o manifests/roomsocket.yaml && \
	kubectl kustomize ./frontend/config/overlays/prod/ -o manifests/frontend.yaml


# kubectl kustomize ./backend/characterimage/config/overlays/prod/ -o manifests/characterimage.yaml && \

gen-model-characterimage:
	cd backend/characterimage/config && \
		sqlboiler psql

# gen-model-characterimage:
# 	gen --sqltype=postgres \
# 		  --connstr "postgres://Snowflake0814:sqvXEFbH7fjJr7@localhost:5400/charimage?sslmode=disable" \
# 			--database charimage \
# 			--gorm \
# 			--overwrite \
# 			--out backend/characterimage/postgres

get-kubeapps-secret:
	kubectl get --namespace default secret kubeapps-operator-token -o jsonpath='{.data.token}' -o go-template='{{.data.token | base64decode}}' && echo

storybook:
	cd frontend && \
	yarn storybook
