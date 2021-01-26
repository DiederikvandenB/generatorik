import React from 'react';

{{ if .HasStyles -}}
import { Styled{{.ComponentName}} } from './{{.ComponentName}}.styles';

{{ end -}}

{{ if .HasProps -}}
interface {{.ComponentName}}Props {}

{{ end -}}

const {{.ComponentName}}: React.FC{{if .HasProps }}<{{.ComponentName}}Props>{{end}} = () => (
  <div />
);

export default {{.ComponentName}};