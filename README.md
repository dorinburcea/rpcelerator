# RPC Aggregator

Welcome to rpc.ag, the rpc aggregator tool for DApps! It simplifies node provider integration
by balancing requests across multiple providers through a single endpoint. rpc.ag also enhances observability, improving
the overall management and efficiency of your DApp infrastructure. Empower the community with better control and
observability by using rpc.ag as the homepage of DApp blockchain infrastructure. Big brothers started watching node
providers.

## Architecture
The architecture of the RPC aggregator with load balancing is shown in the following flowchart:

```mermaid
flowchart LR 
    CC(CLI) --> RPCProxy
    RC(Rest) --> RPCProxy
    RPCC(RPC) --> RPCProxy
    WSC(WS) --> RPCProxy
    subgraph RPC Proxy
      RPCProxy(Handler) --> Routing
        subgraph fasthttp
         Routing --> Auth(Auth & Cors) 
         Auth --> Balancer
        end
         Balancer --> Proxy
    end
    Proxy -->|TCP -> 100ms| Upstream1[Provider 1]  
    Proxy -->|TCP -> 101ms| Upstream2[Provider 2] 
    Proxy -->|TCP -> 200ms| Upstream3[Provider 3] 
    Proxy -->|TCP -> 300ms| Upstream4[Provider N]
```

if you are curious about details


## Short-Term TODO

- [x] Implement rate limiting
- [x] Implement CORS Policy and api keys
- [ ] Implement a ping metric fetcher from all providers
- [ ] Implement a healthcheck scheduler to add the provider(node) back when they are functioning properly
- [ ] Detect when a node is behind, remove it and watch closely again
- [ ] Implement prometheus metrics & endpoint
- [ ] Implement an API lists all nodes (except their API keys) their performance and a simple uptime

## Roadmap

### Phase 0: Release v0 ⌛️

- [ ] Distribute api keys for every project attends to [grizzlython](https://solana.com/grizzlython) with ability of;
- [x] Api key auth & CORS
- [x] Rate limiting

### Phase 1: Public Page

- [ ] Build a public page on [rpc.ag](rpc.ag) that showcases the RPC Aggregator and its features.
- [ ] Include information on how to use the RPC Aggregator.
- [ ] Build an uptime page for transparency and accountability.
- [ ] Show fastest & most available node providers with uptime history
- [ ] Show average ping time with ability of filtering by provider and region 

### Phase 2: Prometheus Exporter and Grafana Dashboards

- Build a Prometheus exporter that exposes the collected metrics in a format that can be scraped by Prometheus.
- Create a set of Grafana dashboards that visualize the performance metrics for the RPC Aggregator and its providers.

### Phase 3: Uptime Page

- Develop an uptime page that displays the current status of the RPC Aggregator and the providers.
- Integrate the monitoring system with the uptime page to display real-time health and performance metrics.
- Allow users to subscribe alerts for downtime or performance issues.


### Phase 4: Become the home page of blockchain infrastructure

- Yes.

With this roadmap, we aim to add a robust monitoring and performance tracking system to the RPC Aggregator, along with a
Prometheus exporter, Grafana dashboards, an uptime page, and a public page. These features will provide greater
visibility into the health and performance of the RPC Aggregator and its providers, and promote transparency and
accountability for blockchain world.

## Developers

rpc.ag provides free rpc endpoints for developers and for projects at their super early stage. Please reach us at
info@rpc.ag

## Providers

If you are a node provider and want to support developer community and want to be mentioned here, reach us at
info@rpc.ag

## Contributing

If you would like to contribute to the RPC Aggregator with load balancing, please fork the repository and create a pull
request with your changes. Be sure to include unit tests and adhere to the project's coding style.

Makefile Overview
-----------------

This project uses a Makefile to manage different actions related to the project, including building, running, testing, checking, creating a Docker image, running a Docker container, and cleaning. The following is a brief overview of each target in the Makefile:

-   `build`: builds the project using the `go build` command and generates the binary file in the `bin/` directory
-   `run`: runs the binary file generated by the `build` target and passes the configuration file as an argument
-   `runp`: runs the binary file generated by the `build` target with a private configuration file located in the `_private/` directory
-   `test`: tests the project using the `go test` command
-   `check`: runs various checks on the project, including verifying the module dependencies, building, vetting, and linting the code
-   `image`: builds a Docker image of the project using the `docker build` command
-   `run-docker`: runs a Docker container of the project using the `docker run` command and maps port 8080
-   `clean`: removes the binary file and the `bin/` directory

Usage
-----

1.  Build the project using the `make build` command
2.  Run the project using the `make run` or `make runp` command
3.  Test the project using the `make test` command
4.  Check the project using the `make check` command
5.  Build a Docker image of the project using the `make image` command
6.  Run a Docker container of the project using the `make run-docker` command
7.  Clean the project using the `make clean` command

Note: Some Makefile targets require the `go`, `docker`, `staticcheck`, and `golint` commands to be installed.

> IF, the project get support (grant from any blockchain) to cover some server expenses and a bit more, we will share it
> to developers who contribute to rpc.ag at any level (proxy, doc, monitoring, etc) we will also open some grants for
> huge tasks and/or issues

Total grant so far: $0  
Sol Wallet to support, donate & track `rpcroe9QQug5tfnG5hvZRCv65R27n3JdFpWuxhBkekH`

