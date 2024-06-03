# Rate-Limiter

Experiment with rate limiter in go. This repo is a mono repo, `exampleserver` is a simple http server implementation with echo, `ratelimiter` is the package that exampleserver imports

## Questions

- How should the repository be structured? Monorepo?
- How should we store data in redis? Problems with using redis? As the number of users scale up do we face any challenges?
- How should we handle the rules?
- What happens when the client is behind a proxy/VPN server?
