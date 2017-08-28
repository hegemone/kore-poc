module Kore
  module Comm
    class PlatformAdapter < Kore::Machinery::Base
      attr_accessor :engine
      TRIGGER_PREFIX = '!'

      def message_received(identity, raw)
        self.log.debug "PlatformAdapter#message_received"
        self.log.debug "  identity: #{identity}"
        self.log.debug "  raw: #{raw}"

        if self.is_cmd(raw)
          self.log.debug "Command matched, parsing..."
          begin
            m = self.parse(raw)
            m.originator = {
              identity: identity,
              platform: self.name
            }
          rescue Exception => e
            self.log.warn "Error occurred while parsing message"
            self.log.warn raw
            self.log.warn e.message
          else
            @engine.route_ingress(m)
          end
        end
      end

      #############################################################
      # Interface methods to be overriden
      #############################################################
      def send_message(emsg)
        raise "platform adapter #{self.name} must implement \"send_message\" method"
      end

      def name
        raise 'platform adapters must implement "name" method'
      end

      # Must be non-blocking!
      def listen
        raise "platform adapter #{self.name} must implement \"listen\" method"
      end
      #############################################################

      def is_cmd(raw)
        !!(Regexp.new("^#{TRIGGER_PREFIX}\\S*($| )").match(raw))
      end

      def parse(raw)
        # TODO: TEST ALL THE THINGS
        # Message structure:
        # Example message types:
        # "!bacon" -- no cmd
        # "!satellite ping" -- no content
        # "!droplet start awesomesauce" -- full
        if raw[0] != TRIGGER_PREFIX
          raise "attempting to parse command without trigger prefix: #{TRIGGER_PREFIX}"
        end

        snipped = raw[1..raw.length]
        content = snipped.clone
        tmp = snipped.split(" ")
        plugin = tmp.shift

        Kore::Comm::IngressMessage.new({
          plugin: plugin,
          content: content,
        })
      end
    end
  end
end
