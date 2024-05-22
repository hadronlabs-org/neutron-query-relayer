import { ManagerModule } from '../../types/Module';
import { DropPuppeteer, DropCore } from '../../generated/contractLib';
import { PuppeteerConfig } from './types/config';
import { Context } from '../../types/Context';
import pino from 'pino';
import { runQueryRelayer, waitBlocks } from '../../utils';

const PuppeteerContractClient = DropPuppeteer.Client;
const CoreContractClient = DropCore.Client;

export class CoreModule implements ManagerModule {
  private puppeteerContractClient?: InstanceType<
    typeof PuppeteerContractClient
  >;
  private coreContractClient?: InstanceType<typeof CoreContractClient>;

  constructor(
    private context: Context,
    private log: pino.Logger,
  ) { }

  private _config: PuppeteerConfig;
  get config(): PuppeteerConfig {
    return this._config;
  }

  init() {
    this.prepareConfig();

    if (this.config.puppeteerContractAddress) {
      this.puppeteerContractClient = new PuppeteerContractClient(
        this.context.neutronSigningClient,
        this.config.puppeteerContractAddress,
      );
    }

    if (this.config.coreContractAddress) {
      this.coreContractClient = new CoreContractClient(
        this.context.neutronSigningClient,
        this.config.coreContractAddress,
      );
    }
  }

  async run(): Promise<void> {
    if (!this.puppeteerContractClient || !this.coreContractClient) {
      this.init();
    }

    const coreContractState =
      await this.coreContractClient.queryContractState();
    const puppeteerResponseReceived =
      !!(await this.coreContractClient.queryLastPuppeteerResponse());

    this.log.debug(
      `Core contract state: ${coreContractState}, response received: ${puppeteerResponseReceived}`,
    );

    // coreContractState === 'transfering' && 
    if (puppeteerResponseReceived) {
      this.log.debug(`Response is received`);

      const queryIds = await this.puppeteerContractClient.queryKVQueryIds();

      this.log.info(`Puppeteer query ids: ${JSON.stringify(queryIds)}`);

      const queryIdsArray = queryIds.map(([queryId]) => queryId.toString());

      this.log.info(
        `Puppeteer query ids plain: ${JSON.stringify(queryIdsArray)}`,
      );

      if (queryIdsArray.length > 0) {
        runQueryRelayer(this.context, this.log, queryIdsArray);

        await waitBlocks(this.context, 3, this.log);

        await this.coreContractClient.tick(
          this.context.neutronWalletAddress,
          1.5,
          undefined,
          [],
        );
      }
    }
  }

  prepareConfig(): void {
    this._config = {
      puppeteerContractAddress:
        process.env.PUPPETEER_CONTRACT_ADDRESS ||
        this.context.factoryContractHandler.factoryState.puppeteer_contract,
      coreContractAddress:
        process.env.CORE_CONTRACT_ADDRESS ||
        this.context.factoryContractHandler.factoryState.core_contract,
    };
  }

  static verifyConfig(log: pino.Logger, skipFactory: boolean): boolean {
    if (skipFactory && !process.env.PUPPETEER_CONTRACT_ADDRESS) {
      log.error('PUPPETEER_CONTRACT_ADDRESS is not provided');
      return false;
    }

    if (skipFactory && !process.env.CORE_CONTRACT_ADDRESS) {
      log.error('CORE_CONTRACT_ADDRESS is not provided');
      return false;
    }

    return true;
  }
}
