<#---
title: App deploy to production
tag: appdeployproduction
api: post
---
We start by finding which version tag to use

eventually 

nexi-booking provision webdeployproduction 
#>

$appname = "koksmat-hooks"
$imagename = "koksmat-hooks"
$dnsname = "koksmat-hooks.home.nexi-intra.com"
$org = "koksmat-com"
$inputFile = join-path  $env:KITCHENROOT $appname ".koksmat","koksmat.json"
$port = 8080
if (!(Test-Path -Path $inputFile) ) {
   Throw "Cannot find file at expected path: $inputFile"
} 
$json = Get-Content -Path $inputFile | ConvertFrom-Json
$version = "v$($json.version.major).$($json.version.minor).$($json.version.patch).$($json.version.build)"

<#
The we build the deployment file
#>

$image = "ghcr.io/$org/$($imagename)-app:$($version)"

$config = @"
apiVersion: apps/v1
kind: Deployment
metadata:
  name: $appname-app
spec:
  selector:
    matchLabels:
      app: $appname-app
  replicas: 1
  template:
    metadata:
      labels:
        app: $appname-app
    spec: 
      containers:
      - name: $appname-app
        image: $image
        command: [$appname]
        args: ["serve"]               
        ports:
          - containerPort: $port
        env:
        - name: NAME
          value: VALUE
    
        
---
apiVersion: v1
kind: Service
metadata:
  name: $appname-app
  labels:
    app: $appname-app
    service: $appname-app
spec:
  ports:
  - name: http
    port: 8080
    targetPort: $port
  selector:
    app: $appname-app
---    
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: $appname
spec:
  rules:
  - host: $dnsname
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: $appname-app
            port:
              number: 8080
    

"@

write-host "Applying config" -ForegroundColor Green

write-host $config -ForegroundColor Gray

$config |  kubectl apply -f -