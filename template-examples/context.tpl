import React from 'react';
import noop from 'lodash/noop';

export interface {{.ContextName}}Props {}

export const defaultProps: {{.ContextName}}Props = {};

export const {{.ContextName}} = React.createContext<{{.ContextName}}Props>(defaultProps);

export const {{.ContextName}}Provider: React.FC = ({ children }) => (
  <{{.ContextName}}.Provider
    value={{"{{}}"}}
  >
    {children}
  </{{.ContextName}}.Provider>
);
