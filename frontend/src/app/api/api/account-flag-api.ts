/* tslint:disable */
/* eslint-disable */
/**
 * Spire
 * Spire API documentation
 *
 * The version of the OpenAPI document: 3.0
 * Contact: akkadius1@gmail.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
import { Configuration } from '../configuration';
// Some imports not used depending on template conditions
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from '../base';
// @ts-ignore
import { CrudcontrollersBulkFetchByIdsGetRequest } from '../models';
// @ts-ignore
import { ModelsAccountFlag } from '../models';
/**
 * AccountFlagApi - axios parameter creator
 * @export
 */
export const AccountFlagApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates AccountFlag
         * @param {ModelsAccountFlag} accountFlag AccountFlag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createAccountFlag: async (accountFlag: ModelsAccountFlag, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'accountFlag' is not null or undefined
            if (accountFlag === null || accountFlag === undefined) {
                throw new RequiredError('accountFlag','Required parameter accountFlag was null or undefined when calling createAccountFlag.');
            }
            const localVarPath = `/account_flag`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'PUT', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            const queryParameters = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                queryParameters.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.query) {
                queryParameters.set(key, options.query[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(queryParameters)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            const nonString = typeof accountFlag !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(accountFlag !== undefined ? accountFlag : {})
                : (accountFlag || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Deletes AccountFlag
         * @param {number} id pAccid
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteAccountFlag: async (id: number, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling deleteAccountFlag.');
            }
            const localVarPath = `/account_flag/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'DELETE', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            const queryParameters = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                queryParameters.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.query) {
                queryParameters.set(key, options.query[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(queryParameters)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Gets AccountFlag
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getAccountFlag: async (id: number, includes?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling getAccountFlag.');
            }
            const localVarPath = `/account_flag/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            if (includes !== undefined) {
                localVarQueryParameter['includes'] = includes;
            }

            if (select !== undefined) {
                localVarQueryParameter['select'] = select;
            }


    
            const queryParameters = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                queryParameters.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.query) {
                queryParameters.set(key, options.query[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(queryParameters)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Gets AccountFlags in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getAccountFlagsBulk: async (body: CrudcontrollersBulkFetchByIdsGetRequest, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'body' is not null or undefined
            if (body === null || body === undefined) {
                throw new RequiredError('body','Required parameter body was null or undefined when calling getAccountFlagsBulk.');
            }
            const localVarPath = `/account_flags/bulk`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            const queryParameters = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                queryParameters.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.query) {
                queryParameters.set(key, options.query[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(queryParameters)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            const nonString = typeof body !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(body !== undefined ? body : {})
                : (body || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Lists AccountFlags
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listAccountFlags: async (includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/account_flags`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            if (includes !== undefined) {
                localVarQueryParameter['includes'] = includes;
            }

            if (where !== undefined) {
                localVarQueryParameter['where'] = where;
            }

            if (whereOr !== undefined) {
                localVarQueryParameter['whereOr'] = whereOr;
            }

            if (groupBy !== undefined) {
                localVarQueryParameter['groupBy'] = groupBy;
            }

            if (limit !== undefined) {
                localVarQueryParameter['limit'] = limit;
            }

            if (orderBy !== undefined) {
                localVarQueryParameter['orderBy'] = orderBy;
            }

            if (orderDirection !== undefined) {
                localVarQueryParameter['orderDirection'] = orderDirection;
            }

            if (select !== undefined) {
                localVarQueryParameter['select'] = select;
            }


    
            const queryParameters = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                queryParameters.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.query) {
                queryParameters.set(key, options.query[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(queryParameters)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Updates AccountFlag
         * @param {number} id Id
         * @param {ModelsAccountFlag} accountFlag AccountFlag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateAccountFlag: async (id: number, accountFlag: ModelsAccountFlag, options: any = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            if (id === null || id === undefined) {
                throw new RequiredError('id','Required parameter id was null or undefined when calling updateAccountFlag.');
            }
            // verify required parameter 'accountFlag' is not null or undefined
            if (accountFlag === null || accountFlag === undefined) {
                throw new RequiredError('accountFlag','Required parameter accountFlag was null or undefined when calling updateAccountFlag.');
            }
            const localVarPath = `/account_flag/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, 'https://example.com');
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'PATCH', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            const queryParameters = new URLSearchParams(localVarUrlObj.search);
            for (const key in localVarQueryParameter) {
                queryParameters.set(key, localVarQueryParameter[key]);
            }
            for (const key in options.query) {
                queryParameters.set(key, options.query[key]);
            }
            localVarUrlObj.search = (new URLSearchParams(queryParameters)).toString();
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            const nonString = typeof accountFlag !== 'string';
            const needsSerialization = nonString && configuration && configuration.isJsonMime
                ? configuration.isJsonMime(localVarRequestOptions.headers['Content-Type'])
                : nonString;
            localVarRequestOptions.data =  needsSerialization
                ? JSON.stringify(accountFlag !== undefined ? accountFlag : {})
                : (accountFlag || "");

            return {
                url: localVarUrlObj.pathname + localVarUrlObj.search + localVarUrlObj.hash,
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * AccountFlagApi - functional programming interface
 * @export
 */
export const AccountFlagApiFp = function(configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates AccountFlag
         * @param {ModelsAccountFlag} accountFlag AccountFlag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createAccountFlag(accountFlag: ModelsAccountFlag, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAccountFlag>>> {
            const localVarAxiosArgs = await AccountFlagApiAxiosParamCreator(configuration).createAccountFlag(accountFlag, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Deletes AccountFlag
         * @param {number} id pAccid
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteAccountFlag(id: number, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await AccountFlagApiAxiosParamCreator(configuration).deleteAccountFlag(id, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets AccountFlag
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getAccountFlag(id: number, includes?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAccountFlag>>> {
            const localVarAxiosArgs = await AccountFlagApiAxiosParamCreator(configuration).getAccountFlag(id, includes, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Gets AccountFlags in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getAccountFlagsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAccountFlag>>> {
            const localVarAxiosArgs = await AccountFlagApiAxiosParamCreator(configuration).getAccountFlagsBulk(body, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Lists AccountFlags
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async listAccountFlags(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAccountFlag>>> {
            const localVarAxiosArgs = await AccountFlagApiAxiosParamCreator(configuration).listAccountFlags(includes, where, whereOr, groupBy, limit, orderBy, orderDirection, select, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
        /**
         * 
         * @summary Updates AccountFlag
         * @param {number} id Id
         * @param {ModelsAccountFlag} accountFlag AccountFlag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async updateAccountFlag(id: number, accountFlag: ModelsAccountFlag, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<ModelsAccountFlag>>> {
            const localVarAxiosArgs = await AccountFlagApiAxiosParamCreator(configuration).updateAccountFlag(id, accountFlag, options);
            return (axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
                const axiosRequestArgs = {...localVarAxiosArgs.options, url: (configuration?.basePath || basePath) + localVarAxiosArgs.url};
                return axios.request(axiosRequestArgs);
            };
        },
    }
};

/**
 * AccountFlagApi - factory interface
 * @export
 */
export const AccountFlagApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    return {
        /**
         * 
         * @summary Creates AccountFlag
         * @param {ModelsAccountFlag} accountFlag AccountFlag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createAccountFlag(accountFlag: ModelsAccountFlag, options?: any): AxiosPromise<Array<ModelsAccountFlag>> {
            return AccountFlagApiFp(configuration).createAccountFlag(accountFlag, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Deletes AccountFlag
         * @param {number} id pAccid
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteAccountFlag(id: number, options?: any): AxiosPromise<string> {
            return AccountFlagApiFp(configuration).deleteAccountFlag(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets AccountFlag
         * @param {number} id Id
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getAccountFlag(id: number, includes?: string, select?: string, options?: any): AxiosPromise<Array<ModelsAccountFlag>> {
            return AccountFlagApiFp(configuration).getAccountFlag(id, includes, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Gets AccountFlags in bulk
         * @param {CrudcontrollersBulkFetchByIdsGetRequest} body body
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getAccountFlagsBulk(body: CrudcontrollersBulkFetchByIdsGetRequest, options?: any): AxiosPromise<Array<ModelsAccountFlag>> {
            return AccountFlagApiFp(configuration).getAccountFlagsBulk(body, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Lists AccountFlags
         * @param {string} [includes] Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
         * @param {string} [where] Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [whereOr] Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
         * @param {string} [groupBy] Group by field. Multiple conditions [.] separated Example: field1.field2
         * @param {string} [limit] Rows to limit in response (Default: 10,000)
         * @param {string} [orderBy] Order by [field]
         * @param {string} [orderDirection] Order by field direction
         * @param {string} [select] Column names [.] separated to fetch specific fields in response
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listAccountFlags(includes?: string, where?: string, whereOr?: string, groupBy?: string, limit?: string, orderBy?: string, orderDirection?: string, select?: string, options?: any): AxiosPromise<Array<ModelsAccountFlag>> {
            return AccountFlagApiFp(configuration).listAccountFlags(includes, where, whereOr, groupBy, limit, orderBy, orderDirection, select, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Updates AccountFlag
         * @param {number} id Id
         * @param {ModelsAccountFlag} accountFlag AccountFlag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        updateAccountFlag(id: number, accountFlag: ModelsAccountFlag, options?: any): AxiosPromise<Array<ModelsAccountFlag>> {
            return AccountFlagApiFp(configuration).updateAccountFlag(id, accountFlag, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * Request parameters for createAccountFlag operation in AccountFlagApi.
 * @export
 * @interface AccountFlagApiCreateAccountFlagRequest
 */
export interface AccountFlagApiCreateAccountFlagRequest {
    /**
     * AccountFlag
     * @type {ModelsAccountFlag}
     * @memberof AccountFlagApiCreateAccountFlag
     */
    readonly accountFlag: ModelsAccountFlag
}

/**
 * Request parameters for deleteAccountFlag operation in AccountFlagApi.
 * @export
 * @interface AccountFlagApiDeleteAccountFlagRequest
 */
export interface AccountFlagApiDeleteAccountFlagRequest {
    /**
     * pAccid
     * @type {number}
     * @memberof AccountFlagApiDeleteAccountFlag
     */
    readonly id: number
}

/**
 * Request parameters for getAccountFlag operation in AccountFlagApi.
 * @export
 * @interface AccountFlagApiGetAccountFlagRequest
 */
export interface AccountFlagApiGetAccountFlagRequest {
    /**
     * Id
     * @type {number}
     * @memberof AccountFlagApiGetAccountFlag
     */
    readonly id: number

    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof AccountFlagApiGetAccountFlag
     */
    readonly includes?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof AccountFlagApiGetAccountFlag
     */
    readonly select?: string
}

/**
 * Request parameters for getAccountFlagsBulk operation in AccountFlagApi.
 * @export
 * @interface AccountFlagApiGetAccountFlagsBulkRequest
 */
export interface AccountFlagApiGetAccountFlagsBulkRequest {
    /**
     * body
     * @type {CrudcontrollersBulkFetchByIdsGetRequest}
     * @memberof AccountFlagApiGetAccountFlagsBulk
     */
    readonly body: CrudcontrollersBulkFetchByIdsGetRequest
}

/**
 * Request parameters for listAccountFlags operation in AccountFlagApi.
 * @export
 * @interface AccountFlagApiListAccountFlagsRequest
 */
export interface AccountFlagApiListAccountFlagsRequest {
    /**
     * Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names 
     * @type {string}
     * @memberof AccountFlagApiListAccountFlags
     */
    readonly includes?: string

    /**
     * Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof AccountFlagApiListAccountFlags
     */
    readonly where?: string

    /**
     * Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2
     * @type {string}
     * @memberof AccountFlagApiListAccountFlags
     */
    readonly whereOr?: string

    /**
     * Group by field. Multiple conditions [.] separated Example: field1.field2
     * @type {string}
     * @memberof AccountFlagApiListAccountFlags
     */
    readonly groupBy?: string

    /**
     * Rows to limit in response (Default: 10,000)
     * @type {string}
     * @memberof AccountFlagApiListAccountFlags
     */
    readonly limit?: string

    /**
     * Order by [field]
     * @type {string}
     * @memberof AccountFlagApiListAccountFlags
     */
    readonly orderBy?: string

    /**
     * Order by field direction
     * @type {string}
     * @memberof AccountFlagApiListAccountFlags
     */
    readonly orderDirection?: string

    /**
     * Column names [.] separated to fetch specific fields in response
     * @type {string}
     * @memberof AccountFlagApiListAccountFlags
     */
    readonly select?: string
}

/**
 * Request parameters for updateAccountFlag operation in AccountFlagApi.
 * @export
 * @interface AccountFlagApiUpdateAccountFlagRequest
 */
export interface AccountFlagApiUpdateAccountFlagRequest {
    /**
     * Id
     * @type {number}
     * @memberof AccountFlagApiUpdateAccountFlag
     */
    readonly id: number

    /**
     * AccountFlag
     * @type {ModelsAccountFlag}
     * @memberof AccountFlagApiUpdateAccountFlag
     */
    readonly accountFlag: ModelsAccountFlag
}

/**
 * AccountFlagApi - object-oriented interface
 * @export
 * @class AccountFlagApi
 * @extends {BaseAPI}
 */
export class AccountFlagApi extends BaseAPI {
    /**
     * 
     * @summary Creates AccountFlag
     * @param {AccountFlagApiCreateAccountFlagRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AccountFlagApi
     */
    public createAccountFlag(requestParameters: AccountFlagApiCreateAccountFlagRequest, options?: any) {
        return AccountFlagApiFp(this.configuration).createAccountFlag(requestParameters.accountFlag, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Deletes AccountFlag
     * @param {AccountFlagApiDeleteAccountFlagRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AccountFlagApi
     */
    public deleteAccountFlag(requestParameters: AccountFlagApiDeleteAccountFlagRequest, options?: any) {
        return AccountFlagApiFp(this.configuration).deleteAccountFlag(requestParameters.id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets AccountFlag
     * @param {AccountFlagApiGetAccountFlagRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AccountFlagApi
     */
    public getAccountFlag(requestParameters: AccountFlagApiGetAccountFlagRequest, options?: any) {
        return AccountFlagApiFp(this.configuration).getAccountFlag(requestParameters.id, requestParameters.includes, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Gets AccountFlags in bulk
     * @param {AccountFlagApiGetAccountFlagsBulkRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AccountFlagApi
     */
    public getAccountFlagsBulk(requestParameters: AccountFlagApiGetAccountFlagsBulkRequest, options?: any) {
        return AccountFlagApiFp(this.configuration).getAccountFlagsBulk(requestParameters.body, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Lists AccountFlags
     * @param {AccountFlagApiListAccountFlagsRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AccountFlagApi
     */
    public listAccountFlags(requestParameters: AccountFlagApiListAccountFlagsRequest = {}, options?: any) {
        return AccountFlagApiFp(this.configuration).listAccountFlags(requestParameters.includes, requestParameters.where, requestParameters.whereOr, requestParameters.groupBy, requestParameters.limit, requestParameters.orderBy, requestParameters.orderDirection, requestParameters.select, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Updates AccountFlag
     * @param {AccountFlagApiUpdateAccountFlagRequest} requestParameters Request parameters.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof AccountFlagApi
     */
    public updateAccountFlag(requestParameters: AccountFlagApiUpdateAccountFlagRequest, options?: any) {
        return AccountFlagApiFp(this.configuration).updateAccountFlag(requestParameters.id, requestParameters.accountFlag, options).then((request) => request(this.axios, this.basePath));
    }
}