export namespace experimental {
	
	export class Community {
	    identifier: string;
	    name: string;
	    discord_url?: string;
	    wiki_url?: string;
	    require_package_listing_approval: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Community(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.identifier = source["identifier"];
	        this.name = source["name"];
	        this.discord_url = source["discord_url"];
	        this.wiki_url = source["wiki_url"];
	        this.require_package_listing_approval = source["require_package_listing_approval"];
	    }
	}

}

