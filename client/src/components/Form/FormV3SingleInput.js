/* @flow */

import React from 'react';
import Style from 'styles';
import { FormInput } from 'components/Form';
import RowLabel from './RowLabel';

type Props = {
  field: string,
  numbersOnly?: bool,
  rowLabel: string,
  valueLabel: { label: string, field: string },
  value: string | number,
  valueError: string,
  onUpdate: (field: string, value: any) => void,
}

class FormV3SingleInput extends React.Component {
  props: Props

  onUpdate = (value: any) => {
    this.props.onUpdate(
      this.props.field,
      this.props.numbersOnly && value !== '' ? Number(value) : value,
    );
  }

  render() {
    return (
      <RowLabel label={this.props.rowLabel} >
        <FormInput
          type={this.props.valueLabel.field}
          floatingLabelText={this.props.valueLabel.label}
          onChange={this.onUpdate}
          value={this.props.value}
          errorText={this.props.valueError}
          numbersOnly={this.props.numbersOnly}
          style={styles.inputWrapper}
          fullWidth
        />
        <div style={styles.inputWrapper} />
        <div style={styles.inputWrapper} />
      </RowLabel>
    );
  }
}

const styles = {
  inputRow: {
    ...Style.Flex.rootRow,
  },
  inputWrapper: {
    marginRight: Style.Dimens.spacing.normal,
    flex: '1 0 0',
  },
};

export default FormV3SingleInput;