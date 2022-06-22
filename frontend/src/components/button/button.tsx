import React from 'react';
import styles from './button.module.css';
import classnames from 'classnames';

export type ButtonVariant =
  'primary' |
  'secondary' |
  'error';

export type ButtonProps = JSX.IntrinsicElements['button'] & {
  variant: ButtonVariant,
}

export function Button(props: ButtonProps) {
  const classes = [styles.button];

  switch (props.variant) {
  case 'primary': classes.push(styles.primary); break;
  case 'secondary': classes.push(styles.secondary); break;
  case 'error': classes.push(styles.error); break;
  }

  const className = classnames(classes);

  return (
    <button
      {...props}
      className={className}
    />
  );
}
