# insta-motion

`insta-motion` is a simple cli tool that provides an easy way to animate
 images for Instagram stories.

## Installation

> [!NOTE]
> [ffmpeg](https://ffmpeg.org/) must be installed in your system.

```bash
npm install -g insta-motion
```

## Usage

With `insta-motion` installed in your system, simply run the cli providing the
 path to the image file.

```bash
insta-motion --image your_image.jpeg --effect zoom-in
```

This will generate a video output in the right resolution for sharing on your Story.
Below is an example of the output generated using the zoom-in effect.

<div align="center">
    <video src="https://github-production-user-asset-6210df.s3.amazonaws.com/444054/326472305-7d01f44b-7a2c-4382-a0de-43ec3db83a3c.mp4?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVCODYLSA53PQK4ZA%2F20240429%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240429T142148Z&X-Amz-Expires=300&X-Amz-Signature=2b96712fc9dd29756ada86acc5bcc850ae0dd032a10d08a428e0d10bb0e9543c&X-Amz-SignedHeaders=host&actor_id=444054&key_id=0&repo_id=793578629" width="100%"></video>
</div>

## Parameters

* `image` Path to the image file.
* `effect` The effect to apply to the image. Available effects: `zoom-in`, `zoom-out`.

## License

[MIT](./LICENSE)
