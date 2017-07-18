// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

/* tslint:disable:ordered-imports variable-name */
import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class Label extends lumi.Resource implements LabelArgs {
    public readonly name: string;
    public color: string;
    public readonly repo?: string;

    public static get(id: lumi.ID): Label {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): Label[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(args: LabelArgs) {
        super();
        if (lumirt.defaultIfComputed(args.name, "") === undefined) {
            throw new Error("Missing required argument 'name'");
        }
        this.name = args.name;
        if (lumirt.defaultIfComputed(args.color, "") === undefined) {
            throw new Error("Missing required argument 'color'");
        }
        this.color = args.color;
        this.repo = args.repo;
    }
}

export interface LabelArgs {
    readonly name: string;
    color: string;
    readonly repo?: string;
}

