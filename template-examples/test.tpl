import React from 'react';

import { render } from 'utils/setupTests';

import {{.ComponentName}} from './{{.ComponentName}}';

describe('<{{.ComponentName}} />', () => {
  it('renders without crashing', () => {
    render(<{{.ComponentName}} />);
    // expect(element).toMatchSnapshot();
  });
});
