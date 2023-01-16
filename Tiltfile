version_settings(
    check_updates=True,
    constraint='>=0.22.2'
) 

v1alpha1.extension_repo(name='tilt-grafana', url='http://github.com/theothertomelliott/tilt-grafana')
v1alpha1.extension(name='tilt-grafana', repo_name='tilt-grafana', repo_path='')
load('ext://tilt-grafana', 'grafana_kubernetes')

load('ext://namespace', 'namespace_yaml')

update_settings(max_parallel_updates=5)

# Command-line flags
config.define_bool("bare",args=False,usage="If set, runs resources as bare processes without kubernetes")
config.define_bool("bots",args=False,usage="If set, run a set of bots to automatically play games")
config.define_bool("disable_telemetry",args=False,usage="Turn off telemetry resources")

args = config.parse()
bots = "bots" in args and args["bots"]
disable_telemetry = "disable_telemetry" in args and args["disable_telemetry"]

local_resource(
    'tests',
    cmd='make testshort',
    deps=['.'],
    ignore=['.output'],
    labels=["test"],
    allow_parallel=True,
)

local_resource(
    'tests (long)',
    cmd='make test',
    trigger_mode = TRIGGER_MODE_MANUAL,
    auto_init = False,
    labels=["test"],
    allow_parallel=True,
)

local_resource(
    'Playwright - Run a Game',
    cmd='npx playwright test tests/playwright/game.spec.js',
    labels=["test"],
    resource_deps=["web"]
)

local_resource(
    'k6 - Load Test ',
    cmd='k6 run ./tests/k6/test.js',
    trigger_mode = TRIGGER_MODE_MANUAL,
    auto_init = False,
    labels=["test"]
)

endpoints = struct(
    jaeger_http="",
)

def telemetry_kubernetes():
    if disable_telemetry:
        return struct(
            jaeger_http="",
        )
    return grafana_kubernetes()

endpoints = telemetry_kubernetes()

lightstep_access_token=""
if os.path.exists("secrets.yaml"):
    secrets = read_yaml("secrets.yaml")
    lightstep_access_token=secrets["lightstep"]["access_token"]

k8s_yaml(namespace_yaml('tictactoe'))

# Load the base Helm chart for all resources
k8s_yaml(helm(
    'charts/tic-tac-toe',
    namespace='tictactoe',
    set=[
        "storage.type=memory",
        "jaeger.http=http://" + endpoints.jaeger_http + "/api/traces",
        ],
))

# Only needed if using mongodb with the stateful set
# k8s_resource("mongodb-standalone", port_forwards="27017:27017")

docker_build('docker.io/tictactoverengineered/base', 'docker/base')

def server(name, port_forwards=[], port="8080", grpcui_port="8081", local_probe=None):
    local_resource(
        name+"-build",
        'make ' + name,
        deps = ["services/" + name, "common"],
        labels=[name]
    )
    docker_build(
        "docker.io/tictactoverengineered/"+name,
        '.build/' + name  + "/",
        dockerfile = "docker/app/Dockerfile",
        live_update = [
            sync('.build/' + name + '/app', '/root/app'),
            run('./restart.sh'),
        ]
    )
    k8s_resource(
        name, 
        port_forwards=port_forwards, 
        resource_deps=[name+'-build'],
        labels=[name]
    )

def space(name,ports=[]):
    local_resource(
        name+"-build",
        'make ' + name,
        deps = ["services/" + name, "common"],
        labels=["space"]
    )
    docker_build(
        "docker.io/tictactoverengineered/"+name,
        '.build/' + name  + "/",
        dockerfile = "docker/app/Dockerfile",
        live_update = [
            sync('.build/' + name + '/app', '/root/app'),
            run('./restart.sh'),
        ]
    )

    k8s_resource(name+'-0-0', resource_deps=[name+'-build'], labels=["space"])
    k8s_resource(name+'-0-1', resource_deps=[name+'-build'], labels=["space"])
    k8s_resource(name+'-0-2', resource_deps=[name+'-build'], labels=["space"])
    k8s_resource(name+'-1-0', resource_deps=[name+'-build'], labels=["space"])
    k8s_resource(name+'-1-1', resource_deps=[name+'-build'], labels=["space"])
    k8s_resource(name+'-1-2', resource_deps=[name+'-build'], labels=["space"])
    k8s_resource(name+'-2-0', resource_deps=[name+'-build'], labels=["space"])
    k8s_resource(name+'-2-1', resource_deps=[name+'-build'], labels=["space"])
    k8s_resource(name+'-2-2', resource_deps=[name+'-build'], labels=["space"])

server("api", port_forwards="8081:8080",port=8081)
server("web", port_forwards="8080:8080",port=8080, local_probe=probe(http_get=http_get_action(8080)))
server("gamerepo", port_forwards=["8082:8080", "8083:8081"], port=8082,grpcui_port=8083)
server("currentturn", port_forwards=["8084:8080", "8085:8081"], port=8084,grpcui_port=8085)
server("grid", port_forwards=["8086:8080", "8087:8081", "2112:2112"], port=8086,grpcui_port=8087)
server("checker", port_forwards=["8088:8080", "8089:8081"], port=8088,grpcui_port=8089)
server("turncontroller", port_forwards=["8090:8080", "8091:8081"], port=8090,grpcui_port=8091)
server("matchmaker", port_forwards=["8092:8080", "8093:8081"], port=8092,grpcui_port=8093)
space("space")

# Bots
if bots or config.tilt_subcommand == "down":
    # Load the Bots Helm chart
    k8s_yaml(helm(
        'charts/tic-tac-toe-bots',
        namespace='tictactoe',
    ))

    server("bot",port_forwards=["2112:2112"], port=2112)
