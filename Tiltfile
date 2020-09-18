docker_build('api-image', '.', 
    dockerfile='build/api/Dockerfile')
k8s_yaml('deploy/api/deploy.yaml')
k8s_resource('api', port_forwards="8081:8080")

docker_build('web-image', '.', 
    dockerfile='build/web/Dockerfile')
k8s_yaml('deploy/web/deploy.yaml')
k8s_resource('web', port_forwards="8080:8080")