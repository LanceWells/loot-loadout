import React from 'react';
import styles from './button.module.css';

export type ButtonProps = JSX.IntrinsicElements['button']

export default function Button(props: ButtonProps) {
  return (
    <button
      {...props}
      className={styles.error}
    />
  );
}
