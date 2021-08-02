# go-grafana

Golang examples with Grafana integration

- Go chi Web Framework link [go-chi](https://github.com/go-chi/chi)
- Documentation link [Grafana](https://grafana.com/docs/grafana/latest/getting-started)
- Documentation link [Prometheus](https://prometheus.io/docs/introduction/overview)

# Run

- `make docker-build`  Build container
- `make docker-up` Run container
- `make docker-down` Stop container

# Endpoints

- `http://localhost:3000/healthcheck` - Healthcheck
- `http://localhost:3000/metrics` - Prometheus metrics
- `http://localhost:5050` - Grafana dashboard

