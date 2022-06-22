import { useRouter } from 'next/router';
import React from 'react';

export default function CharacterEditorView() {
  const router = useRouter();
  const {
    characterID,
  } = router.query;

  return (
    <div>
      <h1>{`This is a placeholder page for character ${characterID}.`}</h1>
      <CharacterEditorPreview />
      <CharacterEditorControls />
    </div>
  );
}

function CharacterEditorPreview() {
  return (
    <div />
  );
}

function CharacterEditorControls() {
  return (
    <div>
      <CharacterEditorButton />
    </div>
  );
}

function CharacterEditorButton() {
  return (
    <div />
  );
}
