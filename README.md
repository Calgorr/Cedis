<h1>Balad's Technical Interview Challenge - Redis-like In-Memory Cache</h1><h2>Introduction</h2><p>This repository is a solution for Balad's technical interview challenge. The goal of the challenge is to design a Redis-like in-memory cache using the Go programming language.</p><h2>Features</h2><p>The solution includes the following features:</p><ul><li>Set: Set the value for a key</li><li>Get: Get the value for a key</li><li>Del: Delete a key-value pair</li><li>Keys Regex: Retrieve keys matching a regular expression</li><li>Multiple Databases: The solution allows for selecting between multiple databases</li></ul><h2>Running the Repository</h2><p>To run the repository, use the following command:</p><pre><div class="bg-black rounded-md mb-4"><div class="flex items-center relative text-gray-200 bg-gray-800 px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span></span><button class="flex ml-auto gap-2"><svg stroke="currentColor" fill="none" stroke-width="2" viewBox="0 0 24 24" stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4" height="1em" width="1em" xmlns="http://www.w3.org/2000/svg"><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"></path><rect x="8" y="2" width="8" height="4" rx="1" ry="1"></rect></svg></button></div><div class="p-4 overflow-y-auto"><code class="!whitespace-pre hljs language-go"><span class="hljs-keyword">go</span> run main.<span class="hljs-keyword">go</span>
</code></div></div></pre><h2>How to Use</h2><p>You can use the provided commands (set, get, del, and keys regex) to interact with the cache.</p><h2>