export class demoMap {
    private m = new Map();
    
     set(k:string, v:number) :Map<string, number> {
            return this.m.set(k, v);
    }

    get(k:string) :number {
            return this.m.get(k);
    }

    getEntity() {
        const keys = this.m.keys();
        for (let key of keys) {
            console.log(`key: ${key}, value: ${this.m.get(key)}`);
        }
        return this.m;
    }

}