{{ define "templates.storage.mongodb.connection" -}}
    {{- if .Values.storage.mongodb.statefulset -}} 
        mongodb://admin:password@database:27017
    {{- else -}}
        {{ .Values.storage.mongodb.connection }} 
    {{- end}}
{{- end }}