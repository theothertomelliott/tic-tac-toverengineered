load('ext://namespace', 'namespace_yaml')

update_settings(max_parallel_updates=1)

test(
    'tests',
    cmd='go test --short ./...',
    deps=['.'],
    ignore=['.output']
)

test(
    'tests (long)',
    cmd='go test ./...',
    trigger_mode = TRIGGER_MODE_MANUAL,
    auto_init = False,
)

local_resource(
    'play',
    cmd='go run ./bot/cmd/onegame http://localhost:8081',
    trigger_mode = TRIGGER_MODE_MANUAL,
    auto_init = False,
)

local_resource(
    'protos', 
    'earthly +protos', 
    deps=[
        'pkg/game/rpcrepository/repo.proto',
        'pkg/game/rpcrepository/earthfile',
        'pkg/turn/rpcturn/current.proto',
        'pkg/turn/rpcturn/controller.proto',
        'pkg/turn/rpcturn/earthfile',
        'pkg/grid/rpcgrid/grid.proto',
        'pkg/grid/rpcgrid/earthfile',
        'pkg/space/rpcspace/space.proto',
        'pkg/space/rpcspace/earthfile'
    ],
    trigger_mode = TRIGGER_MODE_MANUAL,
    auto_init = False,
)

secrets = read_yaml("secrets.yaml")

k8s_yaml(namespace_yaml('tictactoe'))

# Load the base Helm chart for all resources
k8s_yaml(helm(
    'charts/tic-tac-toe',
    namespace='tictactoe',
    set=[
        "honeycomb.api_key=" + secrets["honeycomb"]["api_key"], 
        "honeycomb.dataset=tictactoe-dev",
        "mongodb.statefulset=true",
        ],
))

# Quick mode builds Go binaries locally for rapid turnaround
# This takes advantage of local caching (which is limited with Earthly)
quick_mode = True

def server(name, port_forwards=[]):
    if quick_mode:
        local_resource(
            name+"-build",
            'VERSION=dev ./' + name + '/build.sh',
            deps = [name, "common"],
            ignore = [
                name + '/.output',
                name + '/views',
            ]
        )
    else:
        local_resource(
            name+"-build",
            'earthly ./' + name + '/+build',
            deps = [name, "common"],
            ignore = [
                name + '/.output',
                name + '/views',
            ]
        )
    custom_build(
        "docker.io/tictactoverengineered/"+name,
        'earthly --build-arg IMAGE_REF=$EXPECTED_REF ./' + name + '/+docker',
        [
            './' + name + '/.output',
            './' + name + '/views'    
        ],
        live_update = [
            sync('./' + name + '/.output/app', '/root/app'),
            sync('./' + name + '/views', '/root/views'),
            run('./restart.sh'),
        ]
    )
    k8s_resource(name, port_forwards=port_forwards, resource_deps=[name+'-build'])

server("api", port_forwards="8081:8080")
server("web", port_forwards="8080:8080")
server("gamerepo", port_forwards=["8082:8080", "8083:8081"])
server("currentturn", port_forwards=["8084:8080", "8085:8081"])
server("grid", port_forwards=["8086:8080", "8087:8081"])
server("checker", port_forwards=["8088:8080", "8089:8081"])
server("turncontroller", port_forwards=["8090:8080", "8091:8081"])
server("bot")
server("space")

k8s_resource("mongodb-standalone", port_forwards="27017:27017")