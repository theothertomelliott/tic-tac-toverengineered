local_resource('protos', 'earth +protos', deps=[
    'pkg/game/rpcrepository/repo.proto',
    'pkg/game/rpcrepository/Earthfile',
    'pkg/turn/rpcturn/current.proto',
    'pkg/turn/rpcturn/controller.proto',
    'pkg/turn/rpcturn/Earthfile',
    'pkg/grid/rpcgrid/grid.proto',
    'pkg/grid/rpcgrid/Earthfile'
])

local_resource(
    'binary-build', 
    'earth +binaries', 
    deps=[
        'cmd', 'internal', 'pkg'
    ], 
    resource_deps=['protos']
)

def server(name, port_forwards):
    custom_build(
        name+'-image',
        'earth --build-arg IMAGE_REF=$EXPECTED_REF ./build/' + name + '/+docker',
        ['./.output/'+name],
        live_update = [
            sync('./.output/'+name, '/root/app'),
            run('./restart.sh'),
        ]
    )
    k8s_yaml('deploy/' + name + '/deploy.yaml')
    k8s_resource(name, port_forwards=port_forwards, resource_deps=['binary-build'])

server("api", "8081:8080")
server("web", "8080:8080")
server("gamerepo", ["8082:8080", "8083:8081"])
server("currentturn", ["8084:8080", "8085:8081"])
server("grid",["8086:8080", "8087:8081"])
server("checker",["8088:8080", "8089:8081"])
server("turncontroller",["8090:8080", "8091:8081"])