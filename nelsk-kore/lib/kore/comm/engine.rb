module Kore
  module Comm
    class Engine < Kore::Machinery::Base
      def initialize(plugins, adapters)
        super() # IMPORTANT: Must call super to get deps
        @config = self.config[:engine]
        self.log.debug 'Engine#initialize'
        @plugins = self.load_extensions(plugins)
        @adapters = self.load_extensions(adapters)
      end
      def start
        self.log.debug 'Engine#start'
        EM.run {
          @adapters.each { |_,adapter| adapter.listen }
        }
      end
      def route_ingress(m)
        self.log.info("Engine#route_ingress")
        begin
          @plugins[m.plugin.to_sym].handle_ingress(m)
        rescue Exception => e
          self.log.error 'Something went wrong trying to handle ingress:'
          self.log.error "  message: #{m}"
          self.log.error "  exception: #{e.message}"
        end
      end
      def handle_egress(originator, raw_emsg)
        @adapters[originator[:platform]].send_message Kore::Comm::EgressMessage.new({
          originator: originator,
          content: raw_emsg
        })
      end
      def load_extensions(extensions)
        extensions.reduce({}) do |container, extension|
          extension = Object::const_get(extension).new
          extension.engine = self
          container[extension.name] = extension
          container
        end
      end
    end
  end
end
