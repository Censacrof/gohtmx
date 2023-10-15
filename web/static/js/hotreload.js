async function connectToHotreloadWs() {
  const retryAfter = 250;
  while (true) {
    try {
      const socket = new WebSocket("ws://" + location.host + "/_hotreloadws");

      return await new Promise((resolve, reject) => {
        socket.addEventListener("open", () => {
          console.info("[HR] Connection to hotreload websocket established");
          return resolve(socket);
        });
        socket.addEventListener("error", reject);
      });
    } catch {
      console.error(
        `[HR] Can't connect to hotreload websocket, retring in ${retryAfter}ms`,
      );
      await new Promise((resolve) => setTimeout(resolve, retryAfter));
    }
  }
}

window.addEventListener("load", async () => {
  const socket = await connectToHotreloadWs();

  socket.addEventListener("close", async () => {
    console.warn(
      "[HR] Connection to hotreload websocket was lost, trying to reconnect...",
    );

    await connectToHotreloadWs();
    console.info("[HR] Reloading...");
    location.reload();
  });
});
