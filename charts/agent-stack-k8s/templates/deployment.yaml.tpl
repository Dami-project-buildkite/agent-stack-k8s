apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Release.Name }}
  namespace: {{ .Release.Namespace }}
spec:
  selector:
    matchLabels:
      app: {{ .Release.Name }}
  template:
    metadata:
      labels:
        app: {{ .Release.Name }}
        {{- toYaml $.Values.labels | nindent 8 }}
      annotations:
        checksum/config: {{ include (print $.Template.BasePath "/config.yaml.tpl") . | sha256sum }}
        checksum/secrets: {{ include (print $.Template.BasePath "/secrets.yaml.tpl") . | sha256sum }}
    spec:
      serviceAccountName: {{ .Release.Name }}-controller
      nodeSelector:
{{ toYaml $.Values.nodeSelector | indent 8 }}
      containers:
      - name: controller
        terminationMessagePolicy: FallbackToLogsOnError
        image: {{ .Values.image }}
        env:
        - name: CONFIG
          value: /etc/config.yaml
        envFrom:
          - secretRef:
              name: {{ .Release.Name }}-secrets
        volumeMounts:
          - name: config
            mountPath: /etc/config.yaml
            subPath: config.yaml
        resources:
          {{- toYaml .Values.resources | nindent 10 }}

        securityContext:
          allowPrivilegeEscalation: false
          readOnlyRootFilesystem: true
          runAsNonRoot: true
          capabilities:
            drop:
            - ALL
          seccompProfile:
            type: RuntimeDefault
      volumes:
        - name: config
          configMap:
            name: {{ .Release.Name }}-config
