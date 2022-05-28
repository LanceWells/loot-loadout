export type Coordinates = {
  x: number
  y: number
}

export type Rectangle = {
  topLeft: Coordinates
  botRight: Coordinates
}

export type Positioning = {
  offset: Coordinates;
  rotation: number;
}

export type SpriteImage = Map<Color, Rectangle[]>

export type Color = string;

/**
 * Returns true if the given color is valid.
 * @param c The color to evaluate.
 * @returns True if the color is valid.
 */
export function Validate(c: Color) {
  return /#[0-9A-F]{6}/.test(c);
}

export interface Sprite {
  image: SpriteImage
  positioning: Positioning;
}

export interface SpriteFactory {
  GetSprite(): Sprite
}

/**
 * A factory type used to create sprites that rely on dynamic information.
 */
export class DynamicSpriteFactory implements SpriteFactory {
  private frameMap: Map<Coordinates, Color>;
  private modelMap: Map<Color, Coordinates>;
  private textureMap: Map<Coordinates, Color>;
  private positioning: Positioning;

  /**
   * Constructs a new sprite factory.
   *
   * @param animationMap This is a map from specific pixel coordinates to the hash value of the
   * pixel that we want to map each coordinate to. Each hash value refers to a key in the
   * {@link modelMap}.
   *
   * @param modelMap This is an intermediary map that links the animation to the color of the pixels
   * in the animation. This looks up the pixel color via the {@link textureMap}, which we then use
   * to determine which color to use at any coordinate in the {@link animationMap}.
   *
   * @param textureMap This is a map from pixel coordinates to the colors that should be represented
   * at the parts represented by the same coordinates in the {@link modelMap}.
   */
  constructor(
    animationMap: Map<Coordinates, Color>,
    modelMap: Map<Color, Coordinates>,
    textureMap: Map<Coordinates, Color>,
    positioning: Positioning,
  ) {
    this.frameMap = animationMap;
    this.modelMap = modelMap;
    this.textureMap = textureMap;
    this.positioning = positioning;
  }

  /**
   * Constructs an {@link Sprite} which describes a single layer of the character image.
   */
  GetSprite(): Sprite {
    const i: Sprite = {
      image: new Map<Color, Rectangle[]>(),
      positioning: this.positioning,
    };

    for (const f of this.frameMap) {
      const m = this.modelMap.get(f[1]);
      if (!m) {
        console.warn('could not find a matching model coordinate', f[1]);
        continue;
      }

      const t = this.textureMap.get(m);
      if (!t) {
        console.warn('could not find a matching texture coordinate', m);
        continue;
      }

      if (!i.image.has(t)) {
        i.image.set(t, []);
      }

      i.image.get(t)!.push({
        botRight: f[0],
        topLeft: f[0],
      });
    }

    return i;
  }
}

/**
 * A factory type used to return sprites that do not require dynamic information.
 */
export class StaticSpriteFacotry implements SpriteFactory {
  private image: SpriteImage;
  private positioning: Positioning;

  /**
   * Constructs a new sprite factory.
   */
  constructor(image: SpriteImage, positioning: Positioning) {
    this.image = image;
    this.positioning = positioning;
  }

  /**
   *
   */
  GetSprite(): Sprite {
    return {
      image: this.image,
      positioning: this.positioning,
    };
  }
}
