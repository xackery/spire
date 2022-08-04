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


import { ModelsNpcType } from './models-npc-type';

/**
 * 
 * @export
 * @interface ModelsPet
 */
export interface ModelsPet {
    /**
     * 
     * @type {number}
     * @memberof ModelsPet
     */
    equipmentset?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsPet
     */
    monsterflag?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsPet
     */
    npc_id?: number;
    /**
     * 
     * @type {ModelsNpcType}
     * @memberof ModelsPet
     */
    npc_type?: ModelsNpcType;
    /**
     * 
     * @type {number}
     * @memberof ModelsPet
     */
    petcontrol?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsPet
     */
    petnaming?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsPet
     */
    petpower?: number;
    /**
     * 
     * @type {number}
     * @memberof ModelsPet
     */
    temp?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelsPet
     */
    type?: string;
}


