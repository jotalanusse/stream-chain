#!/bin/sh

npm run webpack
cp __native__/__ios__/v4-native-client.js ~/v4-native-ios/klyra/klyraPresenters/klyraPresenters/_Features
cp __native__/__ios__/v4-native-client.js ~/native-android/v4/integration/cosmos/src/main/assets/
