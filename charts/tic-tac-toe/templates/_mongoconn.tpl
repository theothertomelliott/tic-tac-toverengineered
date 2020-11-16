{{ define "templates.mongodb.connection" -}}
    {{- if .Values.mongodb.statefulset -}} 
        mongodb://admin:password@database:27017
    {{- else -}}
        {{ .Values.mongodb.connection }} 
    {{- end}}
{{- end }}