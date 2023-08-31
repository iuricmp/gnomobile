export interface GoBridgeInterface {
  initBridge(): Promise<void>;
  closeBridge(): Promise<void>;
  createDefaultAccount(name: string): Promise<string>;
  exportJsonConfig(): Promise<string>;
  createReply(message: string, _: string | null): Promise<string>;
}
