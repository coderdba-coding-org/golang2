FROM docker.company.com/k8s/kubectl:1.15

RUN apk add --update dumb-init 
ADD k8s-autofixer /

ENTRYPOINT ["dumb-init", "/k8s-autofixer"]
