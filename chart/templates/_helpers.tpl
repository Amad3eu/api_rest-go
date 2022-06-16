{{- define "labels" -}}
chart-name: {{ .Chart.Name | quote }}
release-name: {{ .Release.Name | quote }}
{{- range $key, $val := .Values.Labels }}
{{$key}} : {{ $val | quote}}
{{- end }}
{{- end }}
