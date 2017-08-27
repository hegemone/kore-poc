module Kore
  module Comm
    class Engine < Kore::Machinery::Base
      def initialize
        super # IMPORTANT: Must call super to get deps
        @config = self.config[:engine]
        self.log.debug 'Engine#initialize'
        self.log.debug "Demo config injection in engine: foo -> #{@config[:foo]}"

        # TODO: Iterate some kind of config and dynamically load up the adapters
        # hardcoding for now...
        @adapters = {
          discord: Kore::Extension::DiscordAdapter.new(self),
          irc: Kore::Extension::IRCAdapter.new(self),
        }
      end
      def start
        self.log.debug 'Engine#start'
        EM.run {
          @adapters.each { |_,adapter| adapter.listen }
        }
      end
      def route_ingress(m)
        self.log.info("Routing ingress message:")
        self.log.info(m)
      end
    end
  end
end
