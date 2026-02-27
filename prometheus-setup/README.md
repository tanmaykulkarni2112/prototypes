# Setting up prometheus to monitor itself

## Prometheus Self-Monitoring — Do We Need an Exporter?

No ✅ — when **Prometheus monitors itself**, you do **not** need to write any exporter code.

Prometheus automatically exposes its own metrics endpoint.

By default, Prometheus exposes:

```
http://localhost:9090/metrics
```

If your configuration contains:

```yaml
scrape_configs:
  - job_name: "prometheus"
    static_configs:
      - targets: ["localhost:9090"]
```

This means:

> Prometheus ➜ scrapes ➜ its own `/metrics` endpoint

So no custom exporter is required.

---

## When Do You Need an Exporter?

You need an exporter when monitoring something that **does not natively expose Prometheus metrics**.

### Examples

| What You Want to Monitor                  | Exporter Required?    |
| ----------------------------------------- | --------------------- |
| Prometheus itself                         | ❌ No                  |
| Linux server metrics (CPU, memory, disk)  | ✅ Yes (Node Exporter) |
| Spring Boot app (with Micrometer enabled) | ❌ No                  |
| Database (MySQL/Postgres)                 | ✅ Yes                 |
| Random REST API without instrumentation   | ✅ Yes                 |

---

## How to Verify Self-Monitoring

Open:

```
http://localhost:9090/targets
```

You should see:

```
job="prometheus"
state="UP"
```

If the target state is **UP**, self-monitoring is working correctly.


--

# Setting up prometheus and Node exporter for monitoring (TODO)


Create the docker-compose.yml for creating pulling the images for 

1. Prometheus
2. Node exporter

prometheus.yml defines what metrics to collect and from where

adding node exporter enables bisinilty into real system performance

# Steps

## 1

Start prometheus using its configuration files - prometheus.yml
Launch Node Exporter on your server
Verify Node export is accessible at http://localhost:9100/metrics

## 2

open Prometheus UI
confirm both Prometheus and Node Exporter targets appear and show the status UP
Take screenshot of service being up

## 3 
OPen graph tab in prom
Run atleast one of the PROMQL QUeries
    a. rate(node_cpu_seconds_total{mode="idle"}[5m])
    b. node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes * 100
    c. time() - process_start_time_seconds
