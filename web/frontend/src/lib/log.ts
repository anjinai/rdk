export const rcLogConditionally = (req: unknown) => {
  if (window.rcDebug || localStorage.getItem('rc_debug')) {
    console.log('gRPC call:', req);
  }
};
