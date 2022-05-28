import React from 'react';
import {DynamicSpriteFactory, StaticSpriteLayer, SpriteAnimation} from './types';

type CharacterBuilderProps<TParts> = {
  // layers: (DynamicSpriteLayer<TParts> | StaticSpriteLayer<TParts>)[]
  // animation: SpriteAnimation<TParts>
}

export default function CharacterBuilder<T>(props: CharacterBuilderProps<T>) {
  return (
    <canvas>
    </canvas>
  );
}
