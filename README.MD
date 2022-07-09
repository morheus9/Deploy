## 1. Add in your gitlab ci\cd variables for gitlab:

```
 - CI_REGISTRY - must contain docker.io.
 - CI_REGISTRY_IMAGE - must contain index.docker.io/username/image_name.
 - CI_REGISTRY_PASSWORD - must contain your Docker HUB password or your access token.
 - CI_REGISTRY_USER - must contain your Docker HUB username.
```

## 2. Add deploy binary+deploy.sh to any folder.

## 3. Start "deploy" binary (deploy app works on 5000 port) and go to 127.0.0.1:5000 and your app will work.

## 4. For deploy do one of two:

### - Add webhooks on your gitlab with "adress_your_server:5000" without ssl verification. (automatic deployment)

or

### - Go to address on your server "127.0.0.1:5000" and you start process of deploy, it will listen this
### address and start deploy.sh. (semi-automatic deployment)

## 5. Done!