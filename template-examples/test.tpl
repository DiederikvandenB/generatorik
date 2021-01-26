import React from 'react';

import { render } from 'utils/testing';

import {{.ComponentName}} from './{{.ComponentName}}';

describe('<{{.ComponentName}} />', () => {
  it('renders without crashing', () => {
    render(<{{.ComponentName}} />);
    // expect(element).toMatchSnapshot();
  });
});
