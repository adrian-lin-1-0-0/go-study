# ctop
Top-like interface for container metrics
> https://github.com/bcicen/ctop




```
alias ctop='docker run --rm -it \
  --name=ctop-$RANDOM \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /etc/localtime:/etc/localtime:ro \
  quay.io/vektorlab/ctop:latest'

```