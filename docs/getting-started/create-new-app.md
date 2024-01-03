# Create a new App (WIP - not working yet)

## Overview

This guide will walk you through creating a new Gno Mobile App using the
`create-app` command.

## Prerequisites

Check the [Build instruction](../../README.md#build-instructions) guide for
prerequisites.

## Create a new App

To create a new app, run the following command:

```console
cd gnonative (root of the repo)
make new-app APP_NAME=MyApp
```

This will create a new app in the `examples/MyApp` directory containing a basic
integration with Gno.

Once you have created the app, you need to copy some files that are not included
in the current script version:

-   GnoCore Framework - TODO
-   XCode workspace - TODO

## Run the App

To run the app, run the following command:

```console
cd examples/MyApp
yarn start
```

and then run the following command in another terminal:

```console
cd examples/MyApp
npx react-native [run-android|run-ios]
```

## Using Gno in your App

To use Gno in your app, you can import the `useGno` hook from
`@gno/hooks/use-gno`:

```ts
import { useGno } from '@gno/hooks/use-gno';

export default function App() {
  const gno = useGno();

  React.useEffect(() => {
    gno.getRemote()
    .then(res => console.log(res))
    .catch(err => console.log(err));
  }, []);

...
```
